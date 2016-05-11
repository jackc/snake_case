#include "stdafx.h"

#include "snapper.h"

auto parseProblems(istream &s)->shared_ptr<vector<problem>>
{
	int problemCount;
	s >> problemCount;

	auto problems = make_shared<vector<problem>>();

	for (int i = 0; i < problemCount; i++) {
		problem p;
		s >> p.deviceCount >> p.snapCount;
		problems->push_back(p);
	}

	return problems;
}

auto solve(problem p)->bool
{
	if (p.snapCount == 0) {
		return false;
	}

	unsigned int cycle = 1 << p.deviceCount;
	return (p.snapCount + 1) % cycle == 0;
}