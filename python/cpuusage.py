from subprocess import Popen, PIPE

def send_command(ip_addr, cmd, user, password):
    """Execute commands in remote machine using sshpass."""
    output = None
    txtcmd = " ssh -oStrictHostKeyChecking=no -oLogLevel=error -o"\
             "UserKnownHostsFile=/dev/null -oCheckHostIP=no "
    ssh_cmd = "sshpass -p " + password + txtcmd + user +\
              "@" + ip_addr + " " + cmd
    #print("ssh command=%s", ssh_cmd)
    try:
        phandle = Popen([ssh_cmd], shell=True, stdout=PIPE)
        output = phandle.communicate()
        #print("ssh output=%s", output)
    except BaseException as err:
        print("SSHError=%s", err)
        return ""
    return output

def get_cpu_usage(ip, user, password):
    """Fetch cpu usage for target ip using proc/stat."""
    keys = ['cpu','user','nice','system','idle','iowait','irq','softirq',
            'steal','guest', 'guest_nice']
    tmp_pcpulst = []
    prev_cpu_list = []
    cpu_cmd = "sudo cat /proc/stat |grep cpu[0-9]"
    outputbuf = send_command(ip, cpu_cmd, user, password)
    #print (type(outputbuf))
    lines = outputbuf[0].split('\n')
    for values in lines:
        #print("values=", values.split())
        #prev_cpu_list.append(dict(zip(keys,values.split())))
        tmp_pcpulst = dict(zip(keys,values.split()))
        pidle = tmp_pcpulst['idle'] + tmp_idle_list['iowait']
        pnonidle = tmp_pcpulst['user'] + tmp_idle_list['nice']
       

get_cpu_usage("172.27.30.34", "pod", "pod")

        
    
