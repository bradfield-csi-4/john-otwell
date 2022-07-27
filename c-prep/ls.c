#include <dirent.h>
#include <stdio.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>

int mkstringsize(char *name, char *size) {
  struct stat stbuf;
  stat(name, &stbuf);
  sprintf(size, "%li", stbuf.st_size);
  return 0;
}

int printif(int flag_thrown, char *toprint) {
  if (flag_thrown || toprint[0] != '.') {
    char size[50];
    mkstringsize(toprint, size);
    printf("%-24s%10s\n", toprint, size);
  }
  return 0;
}

int dodir(char *dirname) {
  DIR *dirp = opendir(dirname);

  printf("%-24s%10s\n", "name", "size");
  for (struct dirent *entry = readdir(dirp); entry != NULL;
       entry = readdir(dirp)) {
    printif(0, entry->d_name);
  }

  closedir(dirp);
  printf("\n");
  return 0;
}

int domultiple(int ndir, char **dirnames) {
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
    domultiple(argc, argv);
  }
}
