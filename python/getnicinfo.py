import os
import sys
import paramiko
import select

def issue_ssh_cmd( addr, cmd, user, password, blocked=True):
    """@The method implements ssh command for DEVICE_TYPES."""
    def stripped(stripbuf):
        """@The method implements ssh command for DEVICE_TYPES."""
        return "".join(i for i in stripbuf if i not in ['\x08'])
    try:
        ssh = paramiko.SSHClient()
        ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
        ssh.connect(hostname=addr, username=user, password=password)
        print('({}) (blocked={}): {}'.format(addr, blocked, cmd))
        stdin, stdout, stderr = ssh.exec_command(cmd)
        #print("log:{}:{}:{}".format(stdin, stdout, stderr))
        if blocked:
            data = stdout.readlines()
        else:
            data = []
            while not stdout.channel.exit_status_ready():
                # Only print data if there is data to read in the channel
                if stdout.channel.recv_ready():
                    rdl, wtl, xtl = select.select([stdout.channel], [], [],
                                                  0.0)
                    if len(rdl) > 0:
                        # Print data from stdout
                        buf = stdout.channel.recv(1024)
                        buf = stripped(buf)
                        if buf:
                            data.append(buf)
                            #sys.stdout.write(buf)
        ssh.close()
    except paramiko.AuthenticationException:
        print("fail cnct {}:wtl={}xtl={}".format(addr, wtl, xtl))
        data = []
    except BaseException:
        print('SSH Exception for ({}): {}'.format(addr, cmd))
        data = []
    return data



def issue_server_cmd_sudo(addr, cmd, blocked=True):
    """@The method implements ssh command for VMHOST."""
    su_cmd = "echo {} | sudo -S ".format('pod') + cmd
    return issue_ssh_cmd(addr, su_cmd, 'pod', 'pod', blocked)

def get_nic_speed( dev_type, hostip, intf):
    """GET interface Speed."""
    state = 'UNKNOWN'
    speed = 'UNKNOWN'
    try:
        ret = []
        cmd1 = "ethtool {} |grep Speed".format(intf)
        ret = (issue_server_cmd_sudo(hostip,
                                              cmd1, True)[0].strip('\n'))
        if len(ret):
            speed = ret.split(':')[1].strip()
            print("{}:Interface {} Speed={}".
                          format(dev_type, intf, speed))

        cmd2 = "ethtool {} |grep Link".format(intf)
        ret = (issue_server_cmd_sudo(hostip,
                                              cmd2, True)[0].strip('\n'))
        if len(ret):
            state = 'UP' if ret.split(':')[1].strip() == 'yes' else 'DOWN'
            print("{}:Interface {} State={}".
                          format(dev_type, intf, state))
    except BaseException as error:
        print("{}:Failed nic={}ip={}err={}".
                       format(dev_type, intf, hostip, error))
    return(state, speed)

print (get_nic_speed('vmhost','10.10.1.11','ens3' ))
