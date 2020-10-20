#!/bin/bash
ATTEND_ENV="dev"
readonly SCRIPT_DIR=`cd $(dirname $0); pwd`
readonly FILE_NAME="/.env"

usage() {
cat <<_EOT_
Usage:

Description:
  .envファイルを生成するスクリプト

Options:
  --help, -h  print this.
_EOT_
exit 1
}

write_default_or_input() {
  read -p "${1}" input
  if [ -z "$input" ]; then
    input=$3
  fi
  echo "$2=$input" >> $SCRIPT_DIR$FILE_NAME
}

case ${1} in
    help|--help|-h)
        usage
    ;;
    rm|--delete|-d)
        rm $SCRIPT_DIR$FILE_NAME
        exit 0
    ;;
esac


REVISION=$(git rev-parse --short HEAD)

if [ -e $SCRIPT_DIR$FILE_NAME ]; then

  exit 0
else
  touch $SCRIPT_DIR$FILE_NAME
  read -p "which envirnment mode(default:dev):" input
  if [ -n "$input" ]; then
    ATTEND_ENV=$input
  fi
  echo "ATTEND_ENV=$ATTEND_ENV" >> $SCRIPT_DIR$FILE_NAME

  write_default_or_input "config file path(default: ./config/config.yaml):" CONFIG_PATH "./config/config.yaml"
  write_default_or_input "attend database name(default: attend):" ATTEND_DB_NAME attend
  write_default_or_input "attend database user(default: root):" ATTEND_DB_USER root
  write_default_or_input "attend database password(default: root):" ATTEND_DB_PASSWARD root
  write_default_or_input "google application credential file path(default: /gcp/attend-dev-key.json):" GOOGLE_APPLICATION_CREDENTIALS "/gcp/attend-dev-key.json"
  echo "REVISION=$REVISION" >> $SCRIPT_DIR$FILE_NAME
  exit 0
fi
