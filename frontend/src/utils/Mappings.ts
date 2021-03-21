export const getVisibilityFromNumber = (num: number): string => {
    switch(num) {
        case 0:
            return 'Public';
        case 1:
            return 'Unlisted';
        case 2:
            return 'Private';
    }

    return 'UNKNOWN VISIBILITY';
}