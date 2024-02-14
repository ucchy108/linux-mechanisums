#!/usr/bin/python3

import os, sys

ret = os.fork()
if ret == 0:
  print("Child process: pid={}, Parent process is pid={}".format(os.getpid(), os.getppid()))
  os.execve("/bin/echo", ["echo", "pid={}: Hello World!!".format(os.getpid())], {})
  exit()
elif ret > 0:
  print("Parent process: pid={}, Child process is pid={}".format(os.getpid(), ret))
  os.execve("/bin/echo", ["echo", "Parent pid={}: Hello World!!".format(os.getpid())], {})
  exit()

sys.exit(1)
