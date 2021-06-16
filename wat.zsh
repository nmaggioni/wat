# https://github.com/junegunn/fzf/blob/7191ebb615f5d6ebbf51d598d8ec853a65e2274d/shell/key-bindings.zsh

export WAT_PATH="${0:a:h}/wat"

__asel() {
  local FZF_OPTS="-e"
  local cmd="${WAT_PATH}"
  setopt localoptions pipefail no_aliases 2> /dev/null
  local item
  eval "$cmd" | FZF_DEFAULT_OPTS="--height 40% --reverse --bind=ctrl-z:ignore $FZF_OPTS " fzf -m "$@" | while read item; do
    echo -n "${(q)item} "
  done
  local ret=$?
  echo
  return $ret
}

wat-alias-widget() {
  LBUFFER="${LBUFFER}$(__asel | grep -Po '^.+(?=:)')"
  local ret=$?
  zle reset-prompt
  return $ret
}
zle     -N   wat-alias-widget
bindkey '^A' wat-alias-widget
