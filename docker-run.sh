dir=$(pwd)
docker run -it -v $dir/in/:/app/in -v $dir/out-dec:/app/out-dec -v $dir/out-enc:/app/out-enc -p 4450:4450 go-crypto