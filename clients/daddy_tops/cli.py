from prompt_toolkit import prompt
from prompt_toolkit.history import FileHistory
from prompt_toolkit.auto_suggest import AutoSuggestFromHistory
from prompt_toolkit.completion import WordCompleter

server = "http://localhost:8888"
user = "<username_for_logging>"
dtWords = ['new', 'group', 'action', 'show', 'result']
dtComp = WordCompleter(dtWords)

def handleNew(split_inp):
    print(split_inp)
    return

def handleShow(split_inp):
    print(split_inp)
    return

def handleInput(inp):
    split_inp = inp.split()
    first_term = split_inp[0].lower()
    if first_term == "new":
        handleNew(split_inp)
    elif first_term == "show":
        handleShow(split_inp)
    elif first_term == "exit" or first_term == "quit":
        print("Goodbye...")
        exit()
    return

while True:
    user_input = prompt('DT>', 
                        history=FileHistory('.dt_history'), 
                        auto_suggest=AutoSuggestFromHistory(), 
                        completer=dtComp)
    handleInput(user_input)
