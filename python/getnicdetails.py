import subprocess
from subprocess import PIPE, Popen
import json

NOT_SUP_INTF = ['v', 'b', 'l', 'o', 'd']


def get_ansible_fact_str(vhostip, uname, password, factname):
    """Get specific fact."""
    output = ""
    sub_cmd1 = "ansible all -i " + vhostip + ", -e \"ansible_user={} \
               ansible_ssh_pass={}\"".format(uname, password)
    sub_cmd2 = " -m setup | sed '1 s/^.*|.*=>.*$/{/g' | jq -r '.ansible_facts \
               ." + factname + "'"
    sub_cmd = sub_cmd1 + sub_cmd2
    try:
        phandle = Popen([sub_cmd], shell=True, stdout=PIPE)
        output = phandle.communicate()
    except (subprocess.CalledProcessError, OSError) as err:
        msg = "CMD={}>:{}".format(sub_cmd, err)
        #print(msg)
    output = output[0].decode('utf-8').strip()
    return(output)

fact = 'ansible_interfaces'
mystr = get_ansible_fact_str("10.10.1.2", "pod", "pod", fact)
#print(type(mystr))
#print(mystr)
jsonout = json.loads(mystr)
#print(type(jsonout))
print(jsonout)
all_nics_dict = {'intflist':''}
nics_list = []
idict = {'name' :'', 'mac_address' :'', 'speed' :'', 'state':'',
         'ipaddress' :'', 'netmask' :''}
for tmpnic in jsonout:
    if str(tmpnic[:1]) in NOT_SUP_INTF:
        jsonout.remove(tmpnic)
print("Supported NIC=",jsonout)

for iface in jsonout:
    mynic = []
    if iface[:1] in NOT_SUP_INTF:
        continue
    nicstr = get_ansible_fact_str("10.10.1.2", "pod", "pod", "ansible_{}".
                                  format(iface))
    mynic = json.loads(nicstr)
    idict['name'] = str(mynic['device'])
    idict['mac_address'] = mynic['macaddress']
    if mynic.has_key('ipv4'):
        idict['ipaddress'] = mynic['ipv4']['address']
        idict['netmask'] = mynic['ipv4']['netmask']
    print(idict)
    nics_list.append(idict)
all_nics_dict.update({'intflist': nics_list})
    
#print("Final nics_lsits", nics_list)
print("Final intf dict=%s", all_nics_dict)
