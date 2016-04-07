#include <stdio.h>

int main() {
  int paths = 0;
  for(int i = 0; i < (1<<20); i++) {
    int bitCount = __builtin_popcount(i);
    if(bitCount == 10) {
      paths += bitCount;
    }
  }

  printf("%d\n", paths);
  return 0;
}
