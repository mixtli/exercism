def is_isogram(word):
    found = {}
    for c in word.lower():
        if not word.isalpha():
            continue
        if c in found:
            return False
        found[c] = True
    return True
