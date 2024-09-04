
# ERP 기능 개선 및 확장을 위한 연동형 API 시스템

## 프로젝트 개요
이 프로젝트는 ERP 시스템의 기능을 개선하고 확장하기 위해 설계된 API 시스템입니다. 이 API는 다양한 ERP 기능과 통합 및 연동될 수 있도록 개발되었습니다.

## 주요 기능
- **ERP 기능 연동**: 기존 ERP 시스템과의 연동을 통해 기능을 개선 및 확장.
- **Golang 기반 개발**: Go 언어로 API 시스템을 구현하여 경량성과 고성능을 유지.
- **MS SQL 쿼리 지원**: MS SQL 데이터베이스에 적합한 쿼리 제공.
- **확장성 있는 인터페이스**: 다양한 ERP 및 CMS 시스템과의 확장성 높은 통합 인터페이스 제공.

## 개발 환경
- **언어**: Golang
- **데이터베이스**: MS SQL
- **기타**: 확장성과 성능을 고려한 인터페이스 설계

## 사용법

1. API 시스템을 클론하여 로컬 환경에 설치합니다:
   ```bash
   git clone https://github.com/your-repo/erp-api-system.git
   ```

2. 설정 파일을 준비합니다. 프로젝트 디렉토리에서 `.env_TEMP` 파일을 `.env` 파일로 복사한 후, 아래의 환경 변수를 시스템에 맞게 설정합니다:

   ```bash
   cp .env_TEMP .env
   ```

   `.env` 파일에 설정할 변수:

   ```env
   DB_HOST=       # 데이터베이스 호스트 주소
   DB_PORT=       # 데이터베이스 포트
   DB_USER=       # 데이터베이스 사용자 이름
   DB_PASSWORD=   # 데이터베이스 비밀번호
   DB_NAME=       # 사용할 데이터베이스 이름

   SERVER_PORT=8090  # API 서버가 실행될 포트
   ```

3. API 서버를 실행합니다:
   ```bash
   go run main.go
   ```

## 시스템 아키텍처
이 시스템은 CMS 및 다양한 ERP 시스템과의 통합을 염두에 두고 설계되었으며, 필요에 따라 확장할 수 있는 모듈형 구조를 지향합니다.

## 기여 방법
기여를 원하시는 분들은 [CONTRIBUTING.md](CONTRIBUTING.md) 파일을 참조해주세요.

## 라이선스
이 프로젝트는 [MIT 라이선스](LICENSE)를 따릅니다.

---

Developed by **CMS Lab**

---

# ERP Feature Enhancement and Expansion API System

## Project Overview
This project is an API system designed to improve and expand the functionality of ERP systems. The API was developed to integrate and extend various ERP features.

## Key Features
- **ERP Integration**: Enhances and expands functions through integration with existing ERP systems.
- **Developed with Golang**: Implemented using Go to maintain lightweight and high performance.
- **MS SQL Query Support**: Provides queries optimized for MS SQL databases.
- **Scalable Interface**: Designed with a scalable interface to allow easy integration with various ERP and CMS systems.

## Development Environment
- **Language**: Golang
- **Database**: MS SQL
- **Others**: Designed with performance and scalability in mind

## Usage

1. Clone the API system to your local environment:
   ```bash
   git clone https://github.com/your-repo/erp-api-system.git
   ```

2. Prepare the configuration file. Copy the `.env_TEMP` file to `.env` and set the environment variables as follows:

   ```bash
   cp .env_TEMP .env
   ```

   Set the following variables in the `.env` file:

   ```env
   DB_HOST=       # Database host address
   DB_PORT=       # Database port
   DB_USER=       # Database username
   DB_PASSWORD=   # Database password
   DB_NAME=       # Database name

   SERVER_PORT=8090  # The port on which the API server will run
   ```

3. Run the API server:
   ```bash
   go run main.go
   ```

## System Architecture
The system is designed with modular architecture to ensure easy integration with CMS and various ERP systems, allowing future scalability.

## Contributing
Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file if you wish to contribute.

## License
This project is licensed under the [MIT License](LICENSE).

---

Developed by **CMS Lab**

