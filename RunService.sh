#!/bin/bash
# 撰寫人員: Neil_Hsieh
# 撰寫日期：2022/05/28
# 說明： 啟動Golang的服務
#
# 備註：
#   

# 取 OS 系統
SYSTEM=$(uname)

# 執行專案的目錄
WORK_PATH=""
# 執行各容器，須掛載的資料夾位置
VOLUME_PATH=""
# 當前用戶名稱
WHOAMI=""
# 用戶專用名稱
USER_PATH=""

# 執行 RunService.sh 的目錄(透過readlink 獲取執行腳本的絕對路徑，再透過dirname取出目錄)
if [ "$SYSTEM" = "Linux" ]
then
    WORK_PATH=$(dirname $(readlink -f $0))
    VOLUME_PATH=$(dirname $(readlink -f $0))/../
fi

# For Mac
if [ "$SYSTEM" = "Darwin" ]
then
    WORK_PATH=$(dirname $(greadlink -f $0))
    VOLUME_PATH=$(dirname $(greadlink -f $0))/../
    WHOAMI=$(whoami)
    USER_PATH="/Users/$WHOAMI"
fi

# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# 讀取圖片路徑(預設dev路徑)
IMG="$VOLUME_PATH/images"
# 環境變數
ENV="local"
# swagger path
SWAGGER_PATH="$GOPATH/pkg/mod/github.com/swaggo"


# 本機開發須安裝swagger + 初始化文件
if [ ! -d "$SWAGGER_PATH" ]; then
    echo "===== Swagger not exist, prepare to install ===="
   go get github.com/swaggo/swag/cmd/swag
   go install github.com/swaggo/swag/cmd/swag
fi

cd $WORK_PATH
swag init

## air hot reload command init, only not exist
if ! [ -f "$WORK_PATH/.air.toml" ]; then
echo "$WORK_PATH/.air.toml"
    $GOPATH/bin/air init
fi

#############################
#############################
docker network ls | grep "web_service" >/dev/null 2>&1
    if  [ $? -ne 0 ]; then
        docker network create web_service
    fi

echo "ENV=$ENV">.env
echo "IMG=$IMG">>.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env


# 啟動容器服務
USER_PATH=$USER_PATH docker-compose up -d