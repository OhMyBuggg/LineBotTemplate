# use to wakeup app
from apscheduler.schedulers.blocking import BlockingScheduler
import datetime
import urllib.request

sched = BlockingScheduler()

#@sched.scheduled_job('cron', minute='*/2')
@sched.scheduled_job('cron', day_of_week='0-6', hour='10-23', minute='*/20')
def scheduled_job():

    print('========== APScheduler CRON =========')
    # 馬上讓我們瞧瞧
    print('This job runs every day */20 min during 10 a.m. to 11 a.m.')
    # 利用datetime查詢時間
    print(f'{datetime.datetime.now().ctime()}')
    print('========== APScheduler CRON =========')

    url = "https://ohmybuglinebot.herokuapp.com/"
    conn = urllib.request.urlopen(url)
        
    for key, value in conn.getheaders():
        print(key, value)

sched.start()