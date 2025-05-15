#!/usr/bin/env bash
# filepath: c:\Users\Steven\Desktop\TGolang\wait-for-it.sh

set -e

host="$1"
shift
cmd="$@"

until nc -z "$host"; do
  echo "Esperando a que el servicio $host esté disponible..."
  sleep 2
done

echo "El servicio $host está disponible. Ejecutando el comando..."
exec $cmd