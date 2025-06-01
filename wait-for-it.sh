#!/usr/bin/env bash

set -e

hostport="$1"
shift

host=$(echo "$hostport" | cut -d: -f1)
port=$(echo "$hostport" | cut -d: -f2)

until nc -z "$host" "$port"; do
  echo "Esperando a que el servicio $host:$port esté disponible..."
  sleep 2
done

echo "El servicio $host:$port está disponible. Ejecutando el comando..."
exec "$@"