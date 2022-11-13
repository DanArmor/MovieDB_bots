import schedule
import time
import json
import requests
import os
from urllib.parse import urlparse

data = None
Headers = {}
with open("envs/dev.json", "r") as f:
    data = json.load(f)
BaseURL = data["ServerURL"]
Headers["pass"] = data["pass"]
  
def getBackup():
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
            print("Backup done")
            return
        print("Backup fail")
    else:
        code = "Unknown" if response == None else str(response.status_code)
        print("Backup failed. Code: " + code)
  
schedule.every().day.at("10:00").do(getBackup)
schedule.every().day.at("20:00").do(getBackup)
  
while True:
    schedule.run_pending()
    time.sleep(1)