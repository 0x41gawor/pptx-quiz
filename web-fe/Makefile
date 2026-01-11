APP_NAME=pptx-fe
TAG=latest
PORT=58174

# Host path on Raspberry Pi where LARGE videos may exist
VIDEO_HOST_DIR=/srv/pptx/videos

# Expected video files (must BOTH exist to enable volume mount)
VIDEO_PL=$(VIDEO_HOST_DIR)/intro_pl.mp4
VIDEO_EN=$(VIDEO_HOST_DIR)/intro_en.mp4


run:
	npm run dev -- --host


docker-build:
	docker build -t $(APP_NAME):$(TAG) .


docker-run:
	@echo "Starting $(APP_NAME)..."
	@if [ -f "$(VIDEO_PL)" ] && [ -f "$(VIDEO_EN)" ]; then \
		echo "✔ Large video files found on host. Running container WITH volume mount."; \
		echo "  Mounting $(VIDEO_HOST_DIR) -> /usr/share/nginx/html/media/videos (read-only)"; \
		docker run -d \
		  -p $(PORT):80 \
		  --name $(APP_NAME) \
		  -v $(VIDEO_HOST_DIR):/usr/share/nginx/html/media/videos:ro \
		  $(APP_NAME):$(TAG); \
	else \
		echo "⚠ Large video files NOT found on host."; \
		echo "  Falling back to SMALL videos bundled inside Docker image."; \
		docker run -d \
		  -p $(PORT):80 \
		  --name $(APP_NAME) \
		  $(APP_NAME):$(TAG); \
	fi