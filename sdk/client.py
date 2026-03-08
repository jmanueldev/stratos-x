import requests
import yaml

API = "http://localhost:8080"

def submit_mission(file):
    with open(file) as f:
        data = yaml.safe_load(f)
    r = requests.post(f"{API}/missions", json=data)
    print(r.json())

def list_missions():
    r = requests.get(f"{API}/missions/list")
    print(r.json())
