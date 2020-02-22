#!/usr/bin/python3.5
# klein based rest service, plus access to the server with credentials defined in
# file server-auth.db= > user:password
# noinspection SpellCheckingInspection
import sys
import json
import attr
from klein import route, Klein
from zope.interface import implementer
from twisted.cred.portal import IRealm
from twisted.internet.defer import succeed
from twisted.cred.portal import Portal
from twisted.web.resource import IResource
from twisted.web.guard import HTTPAuthSessionWrapper, BasicCredentialFactory
from twisted.cred.checkers import FilePasswordDB
sys.path.append('logger')  # noqa
assert route

# for Json dump format
INDENT = 2


class RestIntf(object):
    """@This class implements functions of REST SERVER."""

    app = Klein()

    def __init__(self):
        """
        @The Base class Constructor.

           @param1: self The object pointer
        """

    def dump_request(self, request):
        """Rest API handler for dumping recieved request in log file."""
        request_data = "REQUEST METHOD={}, PATH={}, ARGS={}".format(
            request.method, request.path, request.args)
        print(request_data)


    @app.route('/', methods=['GET'])
    def pod_config_get(self, request):
        """Rest API handler for dumping configuration."""
        self.dump_request(request)
        result = {'Data':'Ok'}
        return json.dumps(result, indent=INDENT)


@implementer(IRealm)
@attr.s
class CheckRealm(object):
    """@This class implements function to setup HTTP Authentication."""

    resource = attr.ib()

    def requestAvatar(self, avatarId, mind, *interfaces):
        """Perform the HTTP Authentication Setup."""
        # print("RequestAvatar", avatarId, mind, interfaces)
        return succeed((IResource, self.resource, lambda: None))


def resource():
    """Initialize the HTTP Authentication."""
    resthandle = RestIntf()
    realm = CheckRealm(resource=resthandle.app.resource())
    portal = Portal(realm, [FilePasswordDB('./server-auth.db')])
    credential_factory = BasicCredentialFactory(b"http auth realm")
    return HTTPAuthSessionWrapper(portal, [credential_factory])

site = Site(resource())
reactor.listenTCP("9999", site)
reactor.run()
