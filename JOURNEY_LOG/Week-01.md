# Week 1 — Setup & Terminal Basics (Level 1)

## Objectives
- Install WSL + Ubuntu
- Basic Linux navigation & file ops
- Git installed & SSH configured
- First Go program + build

## Quests
- [ ] Install **WSL** and **Ubuntu 22.04** on Windows
- [ ] Update packages: `sudo apt update && sudo apt upgrade -y`
- [ ] Install tools: `git`, `curl`, `build-essential`
- [ ] Generate SSH key and add to GitHub
- [ ] Install **Go** (>= 1.22) and set PATH
- [ ] Run `go version`
- [ ] Create and run the sample CLI in `PROJECTS/cli-hello-go`
- [ ] Push this repo to GitHub

## Commands (cheatsheet)

### In PowerShell (Windows)
```powershell
wsl --install -d Ubuntu-22.04
```

### In Ubuntu (WSL)
```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y git curl build-essential

# SSH key (press enter for defaults)
ssh-keygen -t ed25519 -C "andraanggrawijaya@users.noreply.github.com"
cat ~/.ssh/id_ed25519.pub
# Copy the printed key → add to GitHub (Settings → SSH and GPG keys)

# Install Go (official tarball)
GO_VER=1.22.6
curl -LO https://go.dev/dl/go${GO_VER}.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go${GO_VER}.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc

go version
```

## Submission (for review)
Paste the outputs:
- `wsl.exe -l -v` (from PowerShell)
- `uname -a`
- `git --version`
- `go version`
- Screenshot/terminal log of running the hello CLI

