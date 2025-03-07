# 사다리 게임 프로젝트

## 프로젝트 개요
- 프로젝트명: ladder-game
- 목적: 사다리 게임 웹 애플리케이션 개발
- 현재 상태: UI 서비스에 React 프론트엔드 통합 완료, Kubernetes 배포 준비 중

## 시스템 아키텍처
프로젝트는 마이크로서비스 아키텍처를 따르고 있으며, 다음과 같은 4개의 주요 컴포넌트로 구성되어 있습니다:

1. **UI 서비스**: React 기반 사용자 인터페이스를 제공하며, 사용자로부터 참가자와 결과를 입력받아 백엔드 서비스에 전달합니다.
2. **Backend1 (Ladder Manager)**: 전체 사다리 게임 프로세스를 관리하는 서비스로, 다른 백엔드 서비스들과 통신하여 결과를 조합합니다.
3. **Backend2 (Ladder Generator)**: 참가자 수에 따라 랜덤한 사다리 구조를 생성하는 서비스입니다.
4. **Backend3 (Result Mapper)**: 생성된 사다리 구조를 기반으로 참가자와 결과를 매핑하는 서비스입니다.

### 아키텍처 다이어그램
```
┌─────────────┐     ┌───────────────────────┐
│             │     │                       │
│   사용자     │◄────►    UI 서비스          │
│             │     │  (React + Go/Gin)     │
└─────────────┘     └───────────┬───────────┘
                                │
                                ▼
                    ┌───────────────────────┐
                    │                       │
                    │   Backend1            │
                    │   (Ladder Manager)    │
                    │                       │
                    └───────────┬───────────┘
                                │
                 ┌──────────────┴──────────────┐
                 │                             │
    ┌────────────▼─────────────┐   ┌───────────▼────────────┐
    │                          │   │                        │
    │   Backend2               │   │   Backend3             │
    │   (Ladder Generator)     │   │   (Result Mapper)      │
    │                          │   │                        │
    └──────────────────────────┘   └────────────────────────┘
```

## 기술 스택
- **백엔드**: Go 1.21 (Gin 프레임워크)
- **프론트엔드**: React 18
- **컨테이너화**: Docker
- **오케스트레이션**: Kubernetes
- **컨테이너 레지스트리**: NHN Cloud Container Registry

## 주요 기능
1. **사다리 게임 생성**: 사용자가 입력한 참가자 목록과 결과 목록을 기반으로 랜덤한 사다리 구조를 생성합니다.
2. **결과 매핑**: 생성된 사다리 구조를 따라 각 참가자가 어떤 결과에 도달하는지 계산합니다.
3. **사용자 친화적 UI**: React 기반의 직관적인 사용자 인터페이스를 제공합니다.

### 데이터 흐름 다이어그램
```
┌──────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  ┌──────────┐        ┌───────────┐        ┌──────────┐       ┌────────┐ │
│  │          │        │           │        │          │       │        │ │
│  │ 참가자    │───────►│ 사다리 생성 │───────►│ 경로 계산 │──────►│ 결과   │ │
│  │ 입력      │        │           │        │          │       │ 표시   │ │
│  │          │        │           │        │          │       │        │ │
│  └──────────┘        └───────────┘        └──────────┘       └────────┘ │
│                                                                          │
│  ┌──────────┐                                                            │
│  │          │                                                            │
│  │ 결과 항목 │                                                            │
│  │ 입력      │                                                            │
│  │          │                                                            │
│  └──────────┘                                                            │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘
```

## 주요 구성 파일
- `/ui/frontend/package.json`: React 애플리케이션 설정
- `/ui/frontend/src/App.js`: 메인 React 컴포넌트
- `/ui/frontend/src/index.js`: 애플리케이션 진입점
- `/ui/main.go`: Go 백엔드 서버 설정
- `/backend1/main.go`: Ladder Manager 서비스
- `/backend2/main.go`: Ladder Generator 서비스
- `/backend3/main.go`: Result Mapper 서비스

## 프로젝트 구조
ladder-game/
├── README.md                 # 프로젝트 메인 문서
├── LICENSE                   # 라이센스 파일
├── go.mod                    # Go 모듈 정의
├── go.sum                    # Go 의존성 체크섬
├── ui/                       # UI 서비스 (프론트엔드 + API 게이트웨이)
│   ├── Dockerfile            # UI 서비스 Docker 이미지 빌드 설정
│   ├── main.go               # UI 서비스 Go 백엔드 진입점
│   └── frontend/             # React 프론트엔드 애플리케이션
│       ├── package.json      # npm 패키지 정의 및 스크립트
│       ├── public/           # 정적 파일 디렉토리
│       │   ├── index.html    # HTML 템플릿
│       │   └── manifest.json # 웹 앱 매니페스트
│       └── src/              # React 소스 코드
│           ├── App.js        # 메인 React 컴포넌트
│           ├── App.css       # 메인 컴포넌트 스타일
│           ├── index.js      # React 앱 진입점
│           ├── index.css     # 글로벌 스타일
│           └── reportWebVitals.js # 성능 측정 유틸리티
├── backend1/                 # Ladder Manager 서비스
│   ├── Dockerfile            # Backend1 Docker 이미지 빌드 설정
│   └── main.go               # Backend1 서비스 진입점
├── backend2/                 # Ladder Generator 서비스
│   ├── Dockerfile            # Backend2 Docker 이미지 빌드 설정
│   └── main.go               # Backend2 서비스 진입점
├── backend3/                 # Result Mapper 서비스
│   ├── Dockerfile            # Backend3 Docker 이미지 빌드 설정
│   └── main.go               # Backend3 서비스 진입점
└── k8s/                      # Kubernetes 배포 설정 디렉토리

## 컨테이너화 및 Kubernetes 배포

### Docker 이미지 빌드 및 푸시

모든 서비스의 Docker 이미지를 빌드하고 NHN Cloud Container Registry에 푸시하려면 다음 스크립트를 실행하세요:

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

### Kubernetes 배포

Kubernetes 클러스터에 애플리케이션을 배포하는 방법은 `k8s` 디렉토리의 [README.md](k8s/README.md)를 참조하세요.

간단한 배포 단계:

1. 네임스페이스 생성:
   ```bash
   kubectl apply -f k8s/namespace.yaml
   ```

2. ConfigMap 배포:
   ```bash
   kubectl apply -f k8s/configmap.yaml
   ```

3. 모든 서비스 배포 (Kustomize 사용):
   ```bash
   kubectl apply -k k8s/
   ```

4. 배포 확인:
   ```bash
   kubectl get pods -n ladder-game
   kubectl get svc -n ladder-game
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

## 개발 환경
- Node.js: v20.18.0
- npm: 10.8.2
- React: ^18.2.0
- Go: 1.24.0

## 알려진 이슈 및 고려사항
1. React Scripts 버전 호환성
   - Node.js v20.18.0 이상에서는 React Scripts 3.0.1과의 OpenSSL 호환성 문제가 발생
   - 해결 방법: `export NODE_OPTIONS=--openssl-legacy-provider && npm start` 명령어로 실행
2. OpenSSL 레거시 제공자 사용에 따른 잠재적 보안 고려사항
3. 프로덕션 빌드 및 배포 테스트 필요

## 추가 개선 사항
- Kubernetes 배포 구성 완성
- 테스트 코드 작성
- 에러 처리 강화
- 프로덕션 빌드 최적화
- 로깅 및 모니터링 추가

## Kubernetes 배포 가이드

### 개요
프로젝트는 마이크로서비스 아키텍처를 따르며, 각 서비스는 독립적인 컨테이너로 패키징되어 Kubernetes 클러스터에 배포됩니다.

### 배포 구성 요소
- **네임스페이스**: `ladder-game` 네임스페이스에 모든 리소스가 배포됩니다.
- **Deployment**: 각 서비스별 Deployment가 정의되어 있어 Pod 복제본을 관리합니다.
- **Service**: 각 서비스에 대한 내부 통신을 위한 Service가 정의되어 있습니다.
- **Ingress**: 외부에서 UI 서비스에 접근하기 위한 Ingress 리소스가 구성되어 있습니다.
- **ConfigMap**: 서비스 간 통신을 위한 환경 변수가 정의되어 있습니다.

### 배포 방법
```bash
# 전체 애플리케이션 배포
cd k8s
kubectl apply -k .

# 개별 서비스 배포
kubectl apply -f k8s/ui-deployment.yaml
kubectl apply -f k8s/backend1-deployment.yaml
kubectl apply -f k8s/backend2-deployment.yaml
kubectl apply -f k8s/backend3-deployment.yaml
```

자세한 배포 가이드는 [k8s/README.md](/k8s/README.md) 파일을 참조하세요.

## 개발 환경 설정

### 프론트엔드 개발
```bash
# React 개발 서버 실행 (기본 방법)
cd ui/frontend
npm install
npm start

# Node.js v20.18.0 이상에서 OpenSSL 호환성 문제 해결을 위한 실행 방법
cd ui/frontend
npm install
export NODE_OPTIONS=--openssl-legacy-provider && npm start
```

### 백엔드 개발
```bash
# Go 백엔드 실행
cd ui
go run main.go
```

### Docker 빌드 및 실행
```bash
# UI 서비스 빌드
cd ui
docker build -t ladder-game-ui .

# UI 서비스 실행
docker run -p 8080:8080 ladder-game-ui