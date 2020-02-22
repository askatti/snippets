import os
from subprocess import Popen, PIPE


def send_command(ip_addr, cmd, user, password):                       
    """Execute commands in remote machine using sshpass."""                 
    output = None                                                           
    txtcmd = " ssh -oStrictHostKeyChecking=no -oLogLevel=error \
               -oUserKnownHostsFile=/dev/null \
               -oCheckHostIP=no "
    ssh_cmd = "sshpass -p " + password + txtcmd + user +\
              "@" + ip_addr + " " + cmd
    print("ssh command=", ssh_cmd)
    try:
        phandle = Popen([ssh_cmd], shell=True, stdout=PIPE)
        output = phandle.communicate()
        #print("ssh output=", output)
    except BaseException as err:
        print("SSHError=%s", err)
        return ""
    return output

#cmd = "top -bn1 | grep load | awk '{printf \"%.2f%%\t\t\\n\", $(NF-2)}'"
cmd = "top -bn1 | grep load | awk '{printf \"%.2f%%\t\t\\n\", $(NF-2)}'"
#cmd2 = "free -h |grep Mem"

ip = raw_input('Enter target IP')
ret = send_command(ip, cmd, "pod", "pod")
print("ret=",ret[0].split())
#print("total=",ret[0].split()[1])
#print("used=",ret[0].split()[2])
#print("avail=",ret[0].split()[3])
#print("total =", ret[0].strip().split(' '))
