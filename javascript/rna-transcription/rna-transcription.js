var dnaTranscriber = function() {
	this.map = {
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U'
	}
	this.toRna = function(strand) {
		var rna = ""
		for(var i = 0; i < strand.length; i++) {
			if(!this.map[strand[i]] ) {
				throw Error('Invalid input');
			}
			rna += this.map[strand[i]];
		}
		return rna;
	}

}

module.exports = dnaTranscriber;
