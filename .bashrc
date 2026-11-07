gcp() {
    git add .
    
    if [ -n "$1" ]; then
        git commit -m "$*"
    else
        read -p "Введите сообщение коммита: " msg
        if [ -z "$msg" ]; then
            echo "Отмена: сообщение коммита не может быть пустым."
            return 1
        fi
        git commit -m "$msg"
    fi

    if [ $? -eq 0 ]; then
        read -p "Выполнить git push? (y/n): " confirm
        case "$confirm" in
            [yY]|[yY][eE][sS]|[дД]|[дД][аА])
                if git rev-parse --abbrev-ref --symbolic-full-name @{u} >/dev/null 2>&1; then
                    git push
                else
                    branch=$(git branch --show-current)
                    echo "Ветка '$branch' не привязана к origin. Выполняется: git push -u origin $branch"
                    git push -u origin "$branch"
                fi
                ;;
            *)
                echo "Push отменен."
                ;;
        esac
    fi
}

gcpOld() {
    git add .
    
    # Если сообщение передано аргументом, используем его; иначе запрашиваем
    if [ -n "$1" ]; then
        git commit -m "$*"
    else
        read -p "Введите сообщение коммита: " msg
        if [ -z "$msg" ]; then
            echo "Отмена: сообщение коммита не может быть пустым."
            return 1
        fi
        git commit -m "$msg"
    fi

    if [ $? -eq 0 ]; then
        read -p "Выполнить git push? (y/n): " confirm
        case "$confirm" in
            [yY]|[yY][eE][sS]|[дД]|[дД][аА])
                git push
                ;;
            *)
                echo "Push отменен."
                ;;
        esac
    fi
}