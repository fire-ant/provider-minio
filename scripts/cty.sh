files=($(ls package/crds))
dest=examples/generated/

mkdir -p ${dest}

for f in ${files[@]}; do
    echo "$f"
    cty generate -c package/crds/$f -o ${dest}
done

