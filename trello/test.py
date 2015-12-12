from trelloSimple import TrelloSimple
import os

trello = TrelloSimple(os.environ['TRELLO_KEY']) #get key at https://trello.com/1/appKey/generate#
b = trello.get(['boards','4d5ea62fd76aa1136000000c'])
print b

