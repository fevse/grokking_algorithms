# Bread-First Search
# Поиск в ширину
from collections import deque

def the_chosen_one(name):
    return len(name) == 7

def bfs(name, graph):
    search_queue = deque()
    search_queue += graph[name]
    searched = []

    while search_queue:
        person = search_queue.popleft()
        if not person in searched:
            if the_chosen_one(person):
                print(f'{person} is the chosen one')
                return True
            else:
                search_queue += graph[person]
                searched.append(person)
    return False

graph = {
    "you" : ["alice", "bob", "claire"],
    "alice" : ["peggy"],
    "bob" : ["thom", "jonny"],
    "claire" : ["anuj", "peggy"],
    "peggy" : [],
    "anuj" : [],
    "thom" : ["boromir"],
    "jonny" : [],
    "boromir" : []
}


bfs('you', graph)
