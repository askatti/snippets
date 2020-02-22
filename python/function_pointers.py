#! /usr/bin/env python
"""
:Description:  How to create and use function pointers
"""
from __future__ import print_function
import os
import sys
import subprocess
from subprocess import Popen, PIPE
import fileinput



class FunctionPointerClass(object):
    """Utility functions for verifying function pointers"""

    def __init__(self):
        """Initialize FunctionPointerClass constructor."""


    @staticmethod
    def isSudoers():
        """Check if /etc/sudoers is sane.

        :return: True if sudoers file is sane
        """
        try:
            with open("/etc/sudoers", "r") as fd_sudoers:
                lines = fd_sudoers.readlines()
            for line in lines:
                if "user ALL=(ALL) NOPASSWD:ALL" in line:
                    return True
        except IOError:
            print("{} File /etc/sudoers Not Found".
                  format(os.path.basename(__file__)))
        return False

    @staticmethod
    def check_ifpipdir_exist():
        """Check if .pip dir exists.

        :return: True if dir exists
        """
        return os.path.isdir('/home/user/.pip')


    # Form a dictionary of function pointers to be invoked sequentially
    methods_list = {'isSudoers': isSudoers.__func__,
                    'check_ifpipdir_exist': check_ifpipdir_exist.__func__}


# main method to call the function pointers
if __name__ == "__main__":
    OBJ= FunctionPointerClass()
    SUB_RESULT = False
    FAILED_FUNC_CALLS = 0

    # call validation function pointers created in the FunctionPointerClass class
    for funcName, method in OBJ.methods_list.items():
        try:

            # pass relevant args for relevant functions is as below

            # this function takes ip address to be validated if
            # present in the vm_configuratin file.
            if funcName == 'sayHello':
                SUB_RESULT = method(sys.argv[1])


            # For all other functions call them withour arg
            else:
                SUB_RESULT = method()

            if not SUB_RESULT:
                print("{0} Error in function ={1}".
                      format(os.path.basename(__file__), funcName))
                FAILED_FUNC_CALLS = FAILED_FUNC_CALLS + 1

        except BaseException:
            FAILED_FUNC_CALLS = FAILED_FUNC_CALLS + 1

            # mark the function names which failed
            os.system("echo {} >/tmp/failedfunction.txt".
                      format(funcName))
