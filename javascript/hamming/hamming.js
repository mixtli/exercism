var Hamming = function() {}

Hamming.prototype.compute = function(str1, str2) {
	distance = 0;
	if(str1.length != str2.length) {
		throw Error("DNA strands must be of equal length.");
	}
	for(i=0; i < str1.length; i++) {
		if(str1[i] != str2[i]) {
			distance++;
		}
	}
	return distance;
}
module.exports = Hamming;
