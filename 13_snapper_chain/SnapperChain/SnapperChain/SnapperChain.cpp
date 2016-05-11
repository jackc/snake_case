// SnapperChain.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"

#include <iostream>

#include "snapper.h"

using namespace std;


int main()
{
	auto problems = parseProblems(cin);

	int caseNum = 0;
	for (auto p : *problems) {
		caseNum++;
		cout << "Case #" << caseNum << ": " << (solve(p) ? "ON" : "OFF") << endl;
	}
    return 0;
}




