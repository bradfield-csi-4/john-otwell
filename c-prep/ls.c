#include <dirent.h>
#include <stdio.h>

int printif(int flag_thrown, char *toprint) {
  if (flag_thrown || toprint[0] != '.') {
    printf("%s\n", toprint);
  }
  return 0;
}

int dodir(char *dirname) {
  DIR *dirp = opendir(dirname);

  for (struct dirent *entry = readdir(dirp); entry != NULL;
       entry = readdir(dirp)) {
    printif(0, entry->d_name);
  }

  closedir(dirp);
  printf("\n");
  return 0;
}

int domore(int ndir, char **dirnames) {
  for (int i = 1; i < ndir; i++) {
    printf("%s:\n", dirnames[i]);
    dodir(dirnames[i]);
  }
  return 0;
}

int main(int argc, char **argv) {
  char *targetDir = argc > 1 ? argv[1] : ".";

  if (argc <= 2) {
    dodir(targetDir);
  } else {
    domore(argc, argv);
  }
}
