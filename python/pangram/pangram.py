def is_pangram(sentence):
    letters = {}
    for c in sentence.lower():
        print("char %s" % c)
        if c >= 'a' and c <= 'z':
            letters[c] = True

    return len(letters.keys()) == 26

