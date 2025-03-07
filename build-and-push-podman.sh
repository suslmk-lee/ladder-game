#!/bin/bash

# 빌드 실패 시 스크립트 종료
set -e

# 레지스트리 정보
REGISTRY="44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry"

# 서비스 목록
SERVICES=("ui" "backend1" "backend2" "backend3")
SERVICE_NAMES=("ladder-game-ui" "ladder-game-manager" "ladder-game-generator" "ladder-game-mapper")

# 현재 디렉토리 저장
CURRENT_DIR=$(pwd)

# 각 서비스 빌드 및 푸시
for i in "${!SERVICES[@]}"; do
  SERVICE=${SERVICES[$i]}
  NAME=${SERVICE_NAMES[$i]}
  
  echo "===== Building ${SERVICE} ====="
  cd "$CURRENT_DIR/${SERVICE}"
  
  # Docker 이미지 빌드
  sudo podman build -t "${REGISTRY}/${NAME}:latest" .
  
  # 빌드 성공 시 푸시
  if [ $? -eq 0 ]; then
    echo "===== Pushing ${SERVICE} ====="
    sudo podman push "${REGISTRY}/${NAME}:latest"
  else
    echo "===== Build failed for ${SERVICE} ====="
    exit 1
  fi
done

echo "===== All services built and pushed successfully ====="
