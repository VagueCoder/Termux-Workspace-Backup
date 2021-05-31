# User Defined Bash Profile .bashrc

export WORKSPACE=~/Workspace

# Vars for Go
export GOPATH=$WORKSPACE/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
export GITHUB=$GOPATH/src/github.com/VagueCoder

# Shared Locations
export SHARED=~/storage/shared/termux-stuff
export SONGS=$SHARED/Songs

# Py pkg spotdl shortcut for Spotify downloads
download()
{
        WORKDIR=${PWD};
        cd $SONGS;
        spotdl $1;
        cd $WORKDIR;
        WORKDIR='';
}

# Py pkg moviepy shortcut for trimming video
alias trim='python $WORKSPACE/py/Trim-Video/trim_video.py'

# Unix shorthands that miss in Termux
alias ll='ls -al'

# SSH Related
alias ssh_start='eval "$(ssh-agent -s)" && ssh-add ~/.ssh/id_rsa_github'

# Backup to GitHub
backup()
{
	ssh_start
	git add .
	git commit -m "Updated Backup"
	git push origin --all
}
