#include <iostream>

using namespace std;

auto solve(int deviceCount, int snapCount)->bool;

int main()
{
	int problemCount;
	cin >> problemCount;

	for (int i = 0; i < problemCount; i++) {
		int deviceCount;
		int snapCount;
		
		cin >> deviceCount >> snapCount;
		cout << "Case #" << (i+1) << ": " << (solve(deviceCount, snapCount) ? "ON" : "OFF") << endl;
	}

    return 0;
}

auto solve(int deviceCount, int snapCount)->bool
{
	if (snapCount == 0) {
		return false;
	}

	unsigned int cycle = 1 << deviceCount;
	return (snapCount + 1) % cycle == 0;
}



