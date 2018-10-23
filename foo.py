#!/usr/bin/python
import sys
import webbrowser

recipient = 'joe.mooney@gmail.com'
subject = 'foobar 123 mysubject'

#with open('body.txt', 'r') as b:
#    body = b.read()

#body="this is a test"
#
#body = body.replace(' ', '%20')
#
#webbrowser.open('mailto:?to=' + recipient + '&subject=' + subject + '&body=' + body, new=1)
#
#sys.exit(0)
import smtplib

#send email from python
#this creates a text email and sends it to an address of your choosing directly from python
def mailsend (toaddress,fromaddress,subject,message):
    header = 'To:' + toaddress + '\n' + 'From:' + fromaddress + '\n' + 'Subject: ' + subject +'\n'
    msg = header + '\n' + message

    ##you need to know your correct smtp server for this to work, mine is in the function below
    #smtpserver = smtplib.SMTP("smtp.cox.net",25)

    print("step1")
    smtpserver = smtplib.SMTP("smtp.cox.net",587)
    print("step2")
    #smtpserver.ehlo()
    smtpserver.starttls()
    #smtpserver.ehlo()
    #smtpserver = smtplib.SMTP("smtp.cox.net",465)
    #smtpserver.startssl()
    print("step3")
    smtpserver.login("mooney_j", "Da01rse1")
    #smtpserver.login("mooney_j@cox.com", "0pE^EjowzCMS")


    print("step4")
    smtpserver.sendmail(fromaddress, toaddress, msg)
    smtpserver.close()
"""    
the below calls the function - you should change the address settings for your purposes, and you can change the subject and message details as well to send the message information you want to send
"""
mailsend (' joe.mooney@gmail.com',' joe.mooney@gmail.com','test from python','this is the test message') 
