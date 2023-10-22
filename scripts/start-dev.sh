#!/bin/bash
export PATH=$GOPATH/bin:$PATH

# Mencari alat 'air' dalam PATH
air_executable=$(command -v air)

# Periksa apakah 'air' ditemukan
if [ -z "$air_executable" ]; then
  echo "Error: 'air' tidak ditemukan dalam PATH."
  exit 1
fi

# Jalankan 'air' dengan file konfigurasi '.air.conf'
"$air_executable" -c .air.conf
