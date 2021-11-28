# linx-backend
**원티드 해커톤 Linx 팀의 백엔드 저장소**

## Environment
- 언어: Go(v1.17)
- 프레임워크: echo(v4)
- ORM: entgo(v0.9.1)
- 데이터베이스: Postgresql(v13.2)
- 개발 환경: Docker(v20.10.8) & Docker-Compose(컨테이너들 정의 및 관리)

## How To Run
```bash
git clone https://Wanted-Linx/linx-backend.git
cd linx-backend
docker-compose up -d --build
```
**Requirements**
- docker 및 docker-compose가 설치되어 있어야 합니다.
- `./linx-backend` 위치에 .env 파일 생성(./.env.sample을 참고)
- `./linx-backend/config` 위치에 local.yaml 파일 생성(./config/sample.yaml 참고)


## Data Flow
![데이터흐름도](https://user-images.githubusercontent.com/44899448/143724741-2df12c1e-837c-4434-ad62-9e3ade6cde94.png)

## System Architecture
![시스템 구조](https://user-images.githubusercontent.com/44899448/143724768-5e97166e-5b5f-4d89-9504-d53b160a3eb4.png)

## Backend Architecture
![백엔드 구조](https://user-images.githubusercontent.com/44899448/143724781-43b0c726-45b0-45cf-9e79-55ab1d9be8e1.png)

## Contributor
- 안태건(atg0831) - atg950831@gmail.com

