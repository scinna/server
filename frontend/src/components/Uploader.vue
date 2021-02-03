<template>
  <div id="upload--background">
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
export default {
  name: "Uploader",
  components: {CustomInput},
  data: function() {
    return {
      Title: '',
      Description: '',
      Visibility: 0,
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