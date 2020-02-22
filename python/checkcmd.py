hostip = "172.27.170.13"
password = "pod"
uname = "pod"
factname = "test"

cmd1 = "ansible all -i " + hostip + ", -e \"ansible_user=" + uname         
cmd2 = " ansible_ssh_pass=" + password + "\" -e \""                       
cmd3 = "ansible_ssh_common_args=\'-o StrictHostKeyChecking=no\'\""     
cmd4 = " -m setup | sed '1"                                                
cmd5 = " s/^.*|.*=>.*$/{/g' | jq -r '.ansible_facts." + factname + "'"  
scmd = cmd1 + cmd2 + cmd3 + cmd4 + cmd5 
print scmd
