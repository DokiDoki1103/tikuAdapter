# 如何部署使用

## 您可以使用我们已经部署好的应用
等待建设中...
## 您可以独立将此程序部署到 linux,windows,macos,docker,k8s 等环境中

- windows 直接从release中下载对应的版本即可，解压后运行即可
- linux 直接从release中下载对应的版本即可，解压后运行即可
- macos 直接从release中下载对应的版本即可，解压后运行即可
- docker 使用 Docker 部署，参考下方 Docker 部署说明
- k8s 建设中...

## Docker 部署

### 构建镜像

```bash
docker build -t tiku-adapter .
```

### 运行容器

#### 基础运行（使用默认配置）

```bash
docker run -d \
  --name tiku-adapter \
  -p 8060:8060 \
  -v ./data:/app/data \
  tiku-adapter
```

#### 推荐方式：使用配置文件指定数据库路径

创建 `config.yaml` 文件并设置数据库路径为 `data/tiku.db`：
```yaml
database:
  path: "data/tiku.db"
```

然后运行容器：
```bash
docker run -d \
  --name tiku-adapter \
  -p 8060:8060 \
  -v ./data:/app/data \
  -v ./config.yaml:/app/config.yaml \
  tiku-adapter
```

### 目录结构说明

容器内的目录结构：
```
/app
├── tiku-adapter        # 程序可执行文件
├── config.yaml         # 配置文件（可选挂载）
└── data                # 数据目录（持久化）
    └── tiku.db         # SQLite 数据库文件
```

### 配置说明

在 `config.yaml` 中可以自定义数据库路径：

```yaml
database:
  path: "data/tiku.db"  # SQLite数据库文件路径
```

**注意事项：**
- Docker 部署时，**强烈建议**在配置文件中设置 `database.path: "data/tiku.db"`，然后挂载 `/app/data` 目录
- 如果不指定 `database.path`，默认使用 `tiku.db`（向后兼容，但不适合 Docker 持久化）
- 将数据库路径设置在 `data` 目录下，可以只挂载该目录，不会影响容器内的程序文件
- 本地部署（非 Docker）保持默认即可，数据库会在程序所在目录创建