import { Button, Flex, LoadingPage } from '@weaveworks/weave-gitops';
import { useCallback, useContext, useState } from 'react';
import useNotifications from '../../../contexts/Notifications';
import Preview from '../../Templates/Form/Partials/Preview';
import { EnterpriseClientContext } from '../../../contexts/EnterpriseClient';
import {
  ClusterAutomation,
  RenderAutomationResponse,
} from '../../../cluster-services/cluster_services.pb';

export const PreviewPRModal = ({
  formData,
  getClusterAutomations,
}: {
  formData: any;
  getClusterAutomations: () => ClusterAutomation[];
}) => {
  const [openPreview, setOpenPreview] = useState(false);
  const [previewLoading, setPreviewLoading] = useState<boolean>(false);
  const [prPreview, setPRPreview] = useState<RenderAutomationResponse | null>(
    null,
  );
  const { setNotifications } = useNotifications();
  const { api } = useContext(EnterpriseClientContext);
  const handlePRPreview = useCallback(() => {
    setPreviewLoading(true);
    return api
      .RenderAutomation({ clusterAutomations: getClusterAutomations() })
      .then(data => {
        setOpenPreview(true);
        setPRPreview(data);
      })
      .catch(err => {
        setNotifications([
          {
            message: { text: err.message },
            severity: 'error',
            display: 'bottom',
          },
        ]);
      })
      .finally(() => setPreviewLoading(false));
  }, [
    api,
    getClusterAutomations,
    setOpenPreview,
    setPRPreview,
    setPreviewLoading,
    setNotifications,
  ]);

  return (
    <>
      {previewLoading ? (
        <LoadingPage className="preview-loading" />
      ) : (
        <Flex end className="preview-cta">
          <Button onClick={() => handlePRPreview()}>PREVIEW PR</Button>
        </Flex>
      )}
      {openPreview && prPreview ? (
        <Preview
          context="policyconfig"
          openPreview={openPreview}
          setOpenPreview={setOpenPreview}
          prPreview={prPreview}
          sourceType={formData.source_type}
        />
      ) : null}
    </>
  );
};