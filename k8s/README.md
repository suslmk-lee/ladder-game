# Kubernetes 배포 가이드

## 사전 준비 사항

1. Docker가 설치되어 있어야 합니다.
2. kubectl이 설치되어 있어야 합니다.
3. NHN Cloud Container Registry 접근 권한이 있어야 합니다.

## 이미지 빌드 및 푸시

모든 서비스의 이미지를 빌드하고 NHN Cloud Container Registry에 푸시하려면 다음 스크립트를 실행하세요:

```bash
./build-and-push.sh
```

또는 각 서비스를 개별적으로 빌드하고 푸시할 수 있습니다:

```bash
# UI 서비스
cd ui
docker build -t 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-ui .
docker push 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-ui

# Ladder Manager 서비스
cd ../backend1
docker build -t 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-manager .
docker push 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-manager

# Ladder Generator 서비스
cd ../backend2
docker build -t 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-generator .
docker push 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-generator

# Result Mapper 서비스
cd ../backend3
docker build -t 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-mapper .
docker push 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/ladder-game-mapper
```

## Kubernetes 배포

### 1. 네임스페이스 생성

```bash
kubectl apply -f k8s/namespace.yaml
```

### 2. ConfigMap 배포

```bash
kubectl apply -f k8s/configmap.yaml
```

### 3. 백엔드 서비스 배포

```bash
kubectl apply -f k8s/backend1-deployment.yaml
kubectl apply -f k8s/backend2-deployment.yaml
kubectl apply -f k8s/backend3-deployment.yaml
```

### 4. UI 서비스 배포

```bash
kubectl apply -f k8s/ui-deployment.yaml
```

### 5. 전체 배포 (Kustomize 사용)

```bash
kubectl apply -k k8s/
```

## 배포 확인

```bash
# 파드 상태 확인
kubectl get pods -n ladder-game

# 서비스 상태 확인
kubectl get svc -n ladder-game

# 인그레스 상태 확인
kubectl get ingress -n ladder-game
```

## 문제 해결

### 파드가 시작되지 않는 경우

```bash
# 파드 상세 정보 확인
kubectl describe pod <pod-name> -n ladder-game

# 파드 로그 확인
kubectl logs <pod-name> -n ladder-game
```

### 서비스 연결 문제

```bash
# 서비스 엔드포인트 확인
kubectl get endpoints -n ladder-game
```

### 이미지 풀 에러

NHN Cloud Container Registry에 로그인되어 있는지 확인하세요:

```bash
docker login 44ce789b-kr1-registry.container.nhncloud.com
```

## 리소스 정리

```bash
# 네임스페이스와 모든 리소스 삭제
kubectl delete namespace ladder-game
