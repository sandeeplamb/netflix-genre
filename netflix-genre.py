#!/usr/bin/env python3.7
# coding: utf-8
"""
    Generate Netflix Genre
    Handy tool to save your time to choose movie genre
"""


import random
import sys

def select_genre_word(filename):
    """
    Select Random Genre from File
    """
    with open(filename, 'r') as fname:
        return random.choice(fname.readlines()).replace('\n', '')

def generate_genre():
    """
    Generate Netflix Movie Genre
    """
    GENRE       = select_genre_word("genre")
    GENRE_WORD  = GENRE
    print("We will watch NetFlix Movie from Genre :", GENRE_WORD)


#############################################
### 
if __name__ == "__main__":
    generate_genre()
else:
    print("Not a Script. Exiting now...")
    sys.exit(2)