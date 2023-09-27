# Git Audit
A tool for auditing a directory for dirty git repos.

## The Why
I've been going down the rabbit hole of 
[Infrastructure as Code](https://www.redhat.com/en/topics/automation/what-is-infrastructure-as-code-iac) 
the last couple years-- although I didn't learn the term until 1-2 years ago.
What started as idempotent (ish) bash scripts became an incomplete collection of 
[ansible playbooks](https://opensource.com/article/18/3/manage-workstation-ansible),
and (at the time of writing) is now fairly settled into [Nix](https://nixos.org/) and 
[Home Manager](https://nixos.wiki/wiki/Home_Manager).  One pain point I've been dealing
with over the years is while my "setup" has become more and more idempotent and streamlined
with backing up and redeploying, my "data" that I generate has lagged behind. The bulk of
important data I have on my machine tends to be version controlled, but whenever I'm moving
platforms, I have to manually check each checked out repository and confirm all changes are synced
with the remote server.  Git Audit is a tool I built for myself to provide visibility into
what version controlled repositories I have checked out on my machine as well as whether or not they
are [dirty](https://git-scm.com/docs/git-status/2.11.4).

