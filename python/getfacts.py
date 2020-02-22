"""VMOST CLASS."""
from __future__ import print_function
import os
import subprocess
import json
from pprint import pprint
HFPATH = "../ansible/inventory"
FACTS_FILE = "facts.json"
NOT_SUP_INTF = ['v', 'b', 'l', 'o', 'd']


class VMHOST(object):
    """VMHOST CLASS."""

    def __init__(self):
        """@The base class constructor."""

    def _generate_vmhost_facts_jsonfile(self, vmhostip):
        """Get facts."""
        jsonfile = None
        sub_cmd1 = "ansible all -i " + HFPATH + "/hosts{}".format(vmhostip)
        sub_cmd2 = " -m setup | sed '1 s/^.*|.*=>.*$/{/g' >/tmp/"
        sub_cmd3 = "{0}{1}".format(vmhostip, FACTS_FILE)
        sub_cmd = sub_cmd1 + sub_cmd2 + sub_cmd3
        try:
            phandle = os.system(sub_cmd)
            if phandle:
                print("Error running Popen Command")
                return False
            jsonfile = "/tmp/{}".format(vmhostip) + FACTS_FILE
        except (subprocess.CalledProcessError, OSError) as err:
            print("CMD={}>:{}".format(sub_cmd, err))

        if os.path.isfile(str(jsonfile)):
            return jsonfile
        else:
            print("{}: File not Found".format(jsonfile))
            return False

    def _get_vmhost_facts(self, jsonfile):
        """Load the host facts json file for parsing."""
        ipa = ''
        intfarray = []
        diskarray = []
        try:
            with open(jsonfile) as ffd:
                facts = json.load(ffd)
        except IOError as ferror:
            print("Error opening file:{}".format(ferror))
            return False
        try:
            os_name = str(facts['ansible_facts']['ansible_lsb']['id'])
            os_ver = str(facts["ansible_facts"]["ansible_lsb"]["release"])
            hname = str(facts["ansible_facts"]["ansible_hostname"])
            pname = str(facts["ansible_facts"]["ansible_product_name"])
            pserial = str(facts["ansible_facts"]["ansible_product_serial"])
            pver = str(facts["ansible_facts"]["ansible_product_version"])
            dev_entry = {'os_name': os_name, 'os_version': os_ver,
                         'hostname': hname, 'product_name': pname,
                         'product_serial': pserial,
                         'product_version': pver}
            iface_list = facts["ansible_facts"]["ansible_interfaces"]
            for iface in iface_list:
                if iface[:1] in NOT_SUP_INTF:
                    continue
                intf = "" + "ansible_{}".format(iface) + ""
                mac = str(facts["ansible_facts"][intf]["macaddress"])
                state = str(facts["ansible_facts"][intf]["active"])
                if "ipv4" in str(facts["ansible_facts"][intf].keys()):
                    ipa = str(facts["ansible_facts"][intf]["ipv4"]["address"])
                intfdict = {'name': str(iface), 'mac_address': mac,
                            'ipaddress': ipa, 'state': state,
                            'speed': ''}
                intfarray.append(intfdict)
            dev_entry.update({'interfaces': intfarray})
            for disk in facts["ansible_facts"]["ansible_devices"].keys():
                size = str(facts["ansible_facts"]["ansible_devices"][disk]
                           ["size"])
                vendor = str(facts["ansible_facts"]["ansible_devices"]
                             [disk]["vendor"])
                diskdict = {'name': str(disk), 'size': size, 'vendor': vendor,
                            'available': ''}
                diskarray.append(diskdict)
            dev_entry.update({'disk_info': diskdict})
            mtotal = str(facts["ansible_facts"]["ansible_memtotal_mb"])
            mfree = str(facts["ansible_facts"]["ansible_memfree_mb"])
            mused = int(mtotal) - int(mfree)
            memorydict = {'total': mtotal, 'free': str(mfree), 'used':
                          str(mused)}
            dev_entry.update({'mem_info': memorydict})
            num_cpu = str(facts["ansible_facts"]["ansible_processor_count"])
            num_cpu_core = str(facts["ansible_facts"]
                               ["ansible_processor_cores"])
            num_cpu_threads_core = str(facts["ansible_facts"][
                "ansible_processor_threads_per_core"])
            num_vcpus = str(facts["ansible_facts"]["ansible_processor_vcpus"])
            cpu_infodict = {'num_cpu': num_cpu, 'num_cpu_core': num_cpu_core,
                            'num_cpu_threads_core': num_cpu_threads_core,
                            'num_vcpus': num_vcpus}
            dev_entry.update({'cpu_info': cpu_infodict})
            dev_entry.update({'fan_info': [{'name': '', 'speed': ''}]})
        except KeyError as error:
            print("Key: {} Not found".format(error))
            return False
        os.system("rm {}".format(jsonfile))
        del facts
        if dev_entry:
            return dev_entry
        else:
            return False

    def gen_facts(self, vmhostip):
        """Generate facsts dictionary."""
        vmhost_factfile = self._generate_vmhost_facts_jsonfile(vmhostip)
        return self._get_vmhost_facts(vmhost_factfile)

TEST_OBJ = VMHOST()
pprint(TEST_OBJ.gen_facts("10.10.1.2"))
