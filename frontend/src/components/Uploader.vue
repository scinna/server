<template>
  <!-- https://www.raymondcamden.com/2019/08/08/drag-and-drop-file-upload-in-vuejs -->
  <!-- No need for dropzone -->
  <div id="upload--background" ref="Background" @click="(e) => closeUploader(e)">
    <form id="upload">
      <CustomInput id="Title" type="text" :label="$t('upload.title')" v-model="Title" required/>
      <CustomInput id="Description" type="text" :label="$t('upload.desc')" v-model="Description"/>

      <div class="visibilities">
        <input type="radio" name="visibility" id="VisibilityPrivate" value="0"/>
        <label for="VisibilityPrivate">{{$t('upload.visibility.private.text')}}</label>
        <input type="radio" name="visibility" id="VisibilityUnlisted" value="1"/>
        <label for="VisibilityUnlisted">{{$t('upload.visibility.unlisted.text')}}</label>
        <input type="radio" name="visibility" id="VisibilityPublic" value="2"/>
        <label for="VisibilityPublic">{{$t('upload.visibility.public.text')}}</label>
      </div>

      <div id="dropzone"></div>

      <CustomInput type="submit" :value="$t('upload.send')"/>
    </form>
  </div>
</template>

<script>
import CustomInput from "@/components/CustomInput";
import Dropzone from 'dropzone';

export default {
  name: "Uploader",
  components: {CustomInput},
  props: {
    hide: { type: Function }
  },
  data: function() {
    return {
      Title: '',
      Description: '',
      Visibility: 0,
    }
  },
  mounted() {
    new Dropzone(".dropzone")
  },
  methods: {
    closeUploader(event) {
      if (event.target !== this.$refs.Background) {
        return
      }

      this.hide();
    }
  }
}
</script>

<style lang="scss" scoped>
@import "../assets/Colors";

@import "../assets/CenteredForm";

#upload--background {
  z-index: 999999;

  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;

  background: rgba(black, .5);
}

#upload {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);

  width: 400px;

  @media (max-width: 400px) {
    width: 100%;
  }

  .visibilities {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

    label {
      margin-top: 0;
    }
  }

}

</style>