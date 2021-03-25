import {useDropzone} from "react-dropzone";
import styles from '../assets/scss/components/dropzone.module.scss';
import {useState} from "react";
import useAsyncEffect from "use-async-effect";
import i18n from "i18n-js";

type IDropzoneProps = {
    onFileSelected: (file: File) => void;
};

export function Dropzone({onFileSelected}: IDropzoneProps) {
    const {acceptedFiles, getRootProps, getInputProps} = useDropzone({
        accept: ['image/jpeg', 'image/png', 'image/gif'],
        maxFiles: 1,
    });
    const [preview, setPreview] = useState<string|null>();

    useAsyncEffect(async () => {
        if (acceptedFiles.length === 0 || acceptedFiles.length > 1) {
            await setPreview(null);
            return
        }

        const reader = new FileReader();
        reader.onload = (e: any) => {
            setPreview(e.target.result);
        }
        reader.readAsDataURL(acceptedFiles[0]);

        onFileSelected(acceptedFiles[0]);
    }, [acceptedFiles])

    return <div {...getRootProps({className: 'dropzone ' + styles.Dropzone})}>
        {
            preview
            &&
           <img src={preview} alt={acceptedFiles.length > 0 ? acceptedFiles[0].name : ''}/>
        }
        <input {...getInputProps()} />
        <p>{i18n.t('dropzone.text')}</p>
    </div>
}