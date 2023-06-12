echo "deleting all resources..."
${KUBECTL} delete managed --all --force
echo "all resources deleted!"