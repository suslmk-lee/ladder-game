# 사다리 게임 프로젝트

## 프로젝트 개요
- 프로젝트명: ladder-game
- 목적: 사다리 게임 웹 애플리케이션 개발
- 현재 상태: UI 서비스에 React 프론트엔드 통합 완료

## 시스템 아키텍처
프로젝트는 마이크로서비스 아키텍처를 따르고 있으며, 다음과 같은 4개의 주요 컴포넌트로 구성되어 있습니다:

1. **UI 서비스**: React 기반 사용자 인터페이스를 제공하며, 사용자로부터 참가자와 결과를 입력받아 백엔드 서비스에 전달합니다.
2. **Backend1 (Ladder Manager)**: 전체 사다리 게임 프로세스를 관리하는 서비스로, 다른 백엔드 서비스들과 통신하여 결과를 조합합니다.
3. **Backend2 (Ladder Generator)**: 참가자 수에 따라 랜덤한 사다리 구조를 생성하는 서비스입니다.
4. **Backend3 (Result Mapper)**: 생성된 사다리 구조를 기반으로 참가자와 결과를 매핑하는 서비스입니다.

## 기술 스택
- **백엔드**: Go 1.24.0 (Gin 프레임워크)
- **프론트엔드**: React 18.2.0
- **컨테이너화**: Docker
- **오케스트레이션**: Kubernetes (계획)

## 주요 기능
1. **사다리 게임 생성**: 사용자가 입력한 참가자 목록과 결과 목록을 기반으로 랜덤한 사다리 구조를 생성합니다.
2. **결과 매핑**: 생성된 사다리 구조를 따라 각 참가자가 어떤 결과에 도달하는지 계산합니다.
3. **사용자 친화적 UI**: React 기반의 직관적인 사용자 인터페이스를 제공합니다.

## 주요 구성 파일
- `/ui/frontend/package.json`: React 애플리케이션 설정
- `/ui/frontend/src/App.js`: 메인 React 컴포넌트
- `/ui/frontend/src/index.js`: 애플리케이션 진입점
- `/ui/main.go`: Go 백엔드 서버 설정
- `/backend1/main.go`: Ladder Manager 서비스
- `/backend2/main.go`: Ladder Generator 서비스
- `/backend3/main.go`: Result Mapper 서비스

## React 프론트엔드 구현
1. **프론트엔드 구조 생성**
   - React 애플리케이션 초기 설정
   - 주요 컴포넌트 구현 (App.js, index.js)
   - 사다리 게임 UI 로직 개발
     - 참가자 및 결과 입력 폼
     - 동적 입력 필드 추가/삭제 기능
     - 입력 검증 로직
     - 게임 결과 표시 화면

2. **개발 환경 설정 및 문제 해결**
   - Node.js v20.18.0에서 발생한 OpenSSL 관련 오류 대응
   - `NODE_OPTIONS=--openssl-legacy-provider` 환경 변수 설정
   - package.json 스크립트 수정

3. **백엔드 통합**
   - Go 백엔드 수정
   - React 빌드 파일 제공 로직 추가
   - 정적 파일 라우팅 설정
   - API 엔드포인트 유지

## 환경 변수 및 설정
- `NODE_OPTIONS=--openssl-legacy-provider`
  - 목적: Node.js 최신 버전에서 이전 React Scripts 호환
  - 적용 위치: package.json의 스크립트 섹션

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
```

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

## 개발 환경
- Node.js: v20.18.0
- npm: 10.8.2
- React: ^18.2.0
- Go: 1.24.0