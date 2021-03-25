import React from "react";
import {FormControl, InputLabel, MenuItem, Select} from "@material-ui/core";
import i18n from "i18n-js";

type Props = {
    disabled: boolean;
    selectedVisibility: number;
    setSelectedVisibility: (event: React.ChangeEvent<unknown>) => void;
}

export function VisibilityDropDown({selectedVisibility, setSelectedVisibility, disabled}: Props) {
    return <Select
                value={selectedVisibility}
                disabled={disabled}
                onChange={setSelectedVisibility}
                fullWidth={true}>
            <MenuItem value={0}>{i18n.t('visibility.public')}</MenuItem>
            <MenuItem value={1}>{i18n.t('visibility.unlisted')}</MenuItem>
            <MenuItem value={2}>{i18n.t('visibility.private')}</MenuItem>
        </Select>;
}