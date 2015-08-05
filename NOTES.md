# Data Model

- Target models
	- Multiple `Region`s
	- Each `Host` belongs to a `Region`
	- Multiple `Port`s for each `Host`
	- Multiple `Scan`s, each of which has one or more `Host`s
- Administrative
	- Multiple `Worker`s


# Thoughts

- Scan:
	- The result of a tool (nmap, nikto, etc.)
	- We want to link this to a Port somehow (probably), but how?  Multiple scans
		can result in the same port, so we have a strange many-to-one relationship
		with scans from the same Host.
	- Might want to split up the concept of a scan against a Host (which results
		in some number of Ports) from the concept of scanning a SINGLE Port for
		information (e.g. Nikto, SSL checking, etc. etc.)
			- NmapScan vs SinglePortScan?
- Port:
	- A given port in a non-filtered state (closed/open)
	- TODO: What if we have too many open ports?
- Worker:
	- A host to SSH to
	- Authentication options (user, key/password, sudo true/false)
  - NOTE: can also use a SOCKS proxy - maybe have a seperate configuration for this?


# Checks

- SSL check
  - Send raw SSL handshake bytes to server, can be done from within Go
- Hostname checks
- Attempt to retrieve hostname from CN (etc.) of any SSL certificate
- Cipher checks (e.g. code from `lsciphers`)
