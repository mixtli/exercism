def to_rna(dna):
    xlate = {
            'G': 'C',
            'C': 'G',
            'T': 'A',
            'A': 'U'
            }
    rna  = ""
    for c in dna:
        if c in xlate.keys():
            rna += xlate[c]
        else:
            return("")
    return rna
