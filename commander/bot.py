import logging
import json
import requests
from telegram import Update
from telegram.ext import ApplicationBuilder, ContextTypes, CommandHandler, MessageHandler, filters
import os
from urllib.parse import urlparse

BaseURL = None
Admins = []
Headers = {}
DestMode = None

logging.basicConfig(
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    level=logging.INFO
)

async def reportNonAuth(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id, text="Неавторизированный администратор!")


def adminValidate(func):
    async def wrapper(update, context):
        if update.effective_chat.id in Admins:
            await func(update, context)
        else:
            await reportNonAuth(update, context)
    return wrapper

def generateHandler(url : str, errMes : str):
    @adminValidate
    async def handler(update: Update, context: ContextTypes.DEFAULT_TYPE):
        response = None
        try:
            response = requests.get(BaseURL + url, headers=Headers)
        except:
            pass
        if response != None and response.status_code == 200:
            data = json.loads(response.text)
            status = data["status"]
            desc = data["desc"]
            mes = f"Status: {status}.\nDesc: {desc}"
            await context.bot.send_message(chat_id=update.effective_chat.id, text=mes)
        else:
            code = "Unknown" if response == None else str(response.status_code)
            await context.bot.send_message(chat_id=update.effective_chat.id, text=errMes + " Code: " + code)
    return handler

@adminValidate
async def runBackup(update: Update, context: ContextTypes.DEFAULT_TYPE):
    response = None
    try:
        response = requests.get(BaseURL + "/bots/run_backup", headers=Headers)
    except:
        pass
    if response != None and response.status_code == 200:
        data = json.loads(response.text)
        status = data["status"]
        desc = data["desc"]
        if(status != "Error"):
            fileName = os.path.basename(urlparse(desc).path)
            r = requests.get(desc, allow_redirects=True, headers=Headers)
            with open("backups/" + fileName, "wb") as f:
                f.write(r.content)
            with open("backups/" + fileName, "rb") as f:
                await context.bot.send_document(chat_id=update.effective_chat.id, document=f, read_timeout=50)
                if(DestMode != "Remote"):
                    os.remove(os.path.join("backups", fileName))
            return
        mes = f"Status: {status}.\nDesc: {desc}"
        await context.bot.send_message(chat_id=update.effective_chat.id, text=mes)
    else:
        code = "Unknown" if response == None else str(response.status_code)
        await context.bot.send_message(chat_id=update.effective_chat.id, text="Ошибка создания резервной копии БД." + " Code: " + code)


@adminValidate
async def unknown(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await context.bot.send_message(chat_id=update.effective_chat.id, text="Sorry, I didn't understand that command.")

if __name__ == '__main__':
    data = None
    with open("envs/dev.json", "r") as f:
        data = json.load(f)
    BaseURL = data["ServerURL"]
    Admins = data["Admins"]
    Headers["pass"] = data["pass"]
    DestMode = data["Mode"]
    application = ApplicationBuilder().token(data["token"]).build()
    
    application.add_handler(CommandHandler('health', generateHandler("/bots/health", "Ошибка подключения к серверу MovieDB.")))
    application.add_handler(CommandHandler('health_sql', generateHandler("/bots/healthSQL", "Ошибка подключения к серверу MySQL.")))
    application.add_handler(CommandHandler('run_server', generateHandler("/bots/run_server", "Ошибка подключения к серверу Worker.")))
    application.add_handler(CommandHandler('stop_server', generateHandler("/bots/stop_server", "Ошибка подключения к серверу Worker.")))
    application.add_handler(CommandHandler('run_backup', runBackup))
    application.add_handler(MessageHandler(filters.COMMAND, unknown))
    
    application.run_polling()