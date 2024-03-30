git add .

while getopts "m:" arg; do

    case $arg in
        m) message=$OPTARG;;
    esac

    git commit -m "$message"

    git push origin dev

    gcloud run deploy

done
