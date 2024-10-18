# add to ~/.bashrc
autoenv() {
    if [ -f ".envrc" ]; then
        source .envrc
    fi
}

cd() {
    builtin cd "$@" && autoenv
}
