#pragma once

#include <iostream>
#include <vector>
#include <memory>

using namespace std;

struct problem {
	int deviceCount;
	int snapCount;
};

auto parseProblems(istream &s)->shared_ptr<vector<problem>>;

auto solve(problem p)->bool;